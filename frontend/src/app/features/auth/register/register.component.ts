import { Component } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { HttpErrorResponse } from '@angular/common/http';
import { AuthService } from '../../../core/services/auth.service';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [FormsModule, RouterLink],
  templateUrl: './register.component.html',
  styleUrl: './register.component.css',
})
export class RegisterComponent {
  username = '';
  password = '';
  confirmPassword = '';
  loading = false;
  error = '';
  success = '';

  constructor(
    private auth: AuthService,
    private router: Router,
  ) {}

  onSubmit(): void {
    this.error = '';
    this.success = '';
    if (!this.username.trim()) {
      this.error = 'Please enter a username.';
      return;
    }
    if (!this.password) {
      this.error = 'Please enter a password.';
      return;
    }
    if (this.password !== this.confirmPassword) {
      this.error = 'Passwords do not match.';
      return;
    }
    this.loading = true;
    this.auth.register({ username: this.username.trim(), password: this.password, role: 'user' }).subscribe({
      next: () => {
        this.success = 'Account created. Redirecting to login...';
        setTimeout(() => this.router.navigate(['/login']), 1500);
      },
      error: (err: HttpErrorResponse) => {
        this.loading = false;
        if (err?.status === 0 || err?.message?.includes('Unknown Error')) {
          this.error = 'Cannot reach the server. Please start the backend (e.g. run the Go server on port 3000) and try again.';
        } else {
          this.error = err?.error?.message ?? err?.message ?? 'Registration failed.';
        }
      },
      complete: () => (this.loading = false),
    });
  }
}
