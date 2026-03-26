import { TestBed } from '@angular/core/testing';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { provideHttpClient } from '@angular/common/http';
import { provideRouter } from '@angular/router';
import { AuthService } from './auth.service';

function buildJwt(payload: Record<string, unknown>): string {
  const base64Payload = btoa(JSON.stringify(payload)).replace(/\+/g, '-').replace(/\//g, '_').replace(/=+$/, '');
  return `x.${base64Payload}.y`;
}

describe('AuthService', () => {
  let service: AuthService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    localStorage.clear();
    TestBed.configureTestingModule({
      providers: [AuthService, provideHttpClient(), provideHttpClientTesting(), provideRouter([])],
    });
    service = TestBed.inject(AuthService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
    localStorage.clear();
  });

  it('should store token and user after successful login', () => {
    const token = buildJwt({ userId: 10, role: 'admin' });

    service.login({ username: 'admin', password: 'admin123' }).subscribe();

    const req = httpMock.expectOne((r) => r.url.endsWith('/login'));
    expect(req.request.method).toBe('POST');
    req.flush({ message: 'ok', token, username: 'admin' });

    expect(service.isLoggedIn()).toBeTrue();
    expect(service.isAdmin()).toBeTrue();
    expect(service.user()?.id).toBe(10);
    expect(localStorage.getItem('authToken')).toBe(token);
  });

  it('should clear auth data on logout', () => {
    localStorage.setItem('authToken', 'token');
    localStorage.setItem('user', JSON.stringify({ id: 1, username: 'u', role: 'user' }));
    service.logout();
    expect(localStorage.getItem('authToken')).toBeNull();
    expect(localStorage.getItem('user')).toBeNull();
  });
});
