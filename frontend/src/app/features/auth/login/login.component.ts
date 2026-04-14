import { Component } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { HttpErrorResponse } from '@angular/common/http';
import { AuthService } from '../../../core/services/auth.service';
import { environment } from '../../../../environments/environment';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [FormsModule, RouterLink],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css',
})
export class LoginComponent {
  username = '';
  password = '';
  loading = false;
  error = '';
  seedMessage = '';

  constructor(
    private auth: AuthService,
    private router: Router,
    private http: HttpClient,
  ) {}

  loadDemoData(force = false): void {
    this.seedMessage = '';
    const url = force ? `${environment.apiUrl}/seed?force=1` : `${environment.apiUrl}/seed`;
    this.http.get<{ message: string; booksAdded?: number }>(url).subscribe({
      next: (res) => {
        this.seedMessage = res.message || 'Demo data loaded! Login: admin / admin123';
        if (res.booksAdded != null) {
          this.seedMessage += ` (${res.booksAdded} books)`;
        }
      },
      error: () => {
        this.seedMessage = 'Start the backend first (port 3000).';
      },
    });
  }

  onSubmit(): void {
    this.error = '';
    if (!this.username.trim() || !this.password.trim()) {
      this.error = 'Please enter username and password.';
      return;
    }
    this.loading = true;
    this.auth.login({ username: this.username.trim(), password: this.password }).subscribe({
      next: () => {
        if (this.auth.isAdmin()) {
          this.router.navigate(['/main/admin/books']);
        } else {
          this.router.navigate(['/main/dashboard']);
        }
      },
      error: (err: HttpErrorResponse) => {
        this.loading = false;
        if (err?.status === 0 || err?.message?.includes('Unknown Error')) {
          this.error = 'Cannot reach the server. Please start the backend (e.g. run the Go server on port 3000) and try again.';
        } else {
          this.error = err?.error?.message ?? err?.message ?? 'Login failed.';
        }
      },
      complete: () => (this.loading = false),
    });
  }
}
