import { Injectable, signal, computed } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { Observable, tap, catchError, of } from 'rxjs';
import { environment } from '../../../environments/environment';
import { User } from '../models/user.model';

export interface LoginRequest {
  username: string;
  password: string;
}

export interface RegisterRequest {
  username: string;
  password: string;
  role?: string;
}

export interface AuthResponse {
  message: string;
  token: string;
  username?: string;
  role?: string;
}

@Injectable({ providedIn: 'root' })
export class AuthService {
  private readonly tokenKey = 'authToken';
  private readonly userKey = 'user';

  private currentUser = signal<User | null>(null);
  private token = signal<string | null>(null);

  user = this.currentUser.asReadonly();
  isLoggedIn = computed(() => !!this.currentUser());
  isAdmin = computed(() => this.currentUser()?.role === 'admin');

  constructor(
    private http: HttpClient,
    private router: Router,
  ) {
    this.initFromStorage();
  }

  private initFromStorage(): void {
    const t = localStorage.getItem(this.tokenKey);
    const u = localStorage.getItem(this.userKey);
    if (t && u) {
      try {
        this.token.set(t);
        this.currentUser.set(JSON.parse(u));
      } catch {
        this.clearAuth();
      }
    }
  }

  login(credentials: LoginRequest): Observable<AuthResponse> {
    return this.http.post<AuthResponse>(`${environment.apiUrl}/login`, credentials).pipe(
      tap((res) => {
        if (res.token) {
          localStorage.setItem(this.tokenKey, res.token);
          const payload = this.parseJwt(res.token);
          const id = payload?.['userId'] ?? payload?.['UserID'] ?? payload?.['userID'];
          const role = payload?.['Role'] ?? payload?.['role'] ?? res.role ?? 'user';
          const user: User = {
            id: typeof id === 'number' || typeof id === 'string' ? id : 0,
            username: res.username ?? credentials.username,
            role: (role as string) as 'user' | 'admin',
          };
          localStorage.setItem(this.userKey, JSON.stringify(user));
          this.token.set(res.token);
          this.currentUser.set(user);
        }
      }),
    );
  }

  register(data: RegisterRequest): Observable<{ message: string }> {
    const body = { ...data, role: data.role ?? 'user' };
    return this.http.post<{ message: string }>(`${environment.apiUrl}/register`, body);
  }

  logout(): void {
    this.clearAuth();
    this.router.navigate(['/login']);
  }

  getToken(): string | null {
    return this.token() ?? localStorage.getItem(this.tokenKey);
  }

  getUserId(): number | string | null {
    return this.currentUser()?.id ?? null;
  }

  private clearAuth(): void {
    localStorage.removeItem(this.tokenKey);
    localStorage.removeItem(this.userKey);
    this.token.set(null);
    this.currentUser.set(null);
  }

  private parseJwt(token: string): Record<string, unknown> | null {
    try {
      const base64Url = token.split('.')[1];
      const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
      const json = decodeURIComponent(
        atob(base64)
          .split('')
          .map((c) => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2))
          .join(''),
      );
      return JSON.parse(json) as Record<string, unknown>;
    } catch {
      return null;
    }
  }
}
