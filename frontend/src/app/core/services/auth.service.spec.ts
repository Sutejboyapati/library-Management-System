import { TestBed } from '@angular/core/testing';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting, HttpTestingController } from '@angular/common/http/testing';
import { provideRouter, Router } from '@angular/router';
import { AuthService } from './auth.service';
import { environment } from '../../../environments/environment';

describe('AuthService', () => {
  let service: AuthService;
  let httpMock: HttpTestingController;
  let router: Router;

  beforeEach(() => {
    localStorage.clear();

    TestBed.configureTestingModule({
      providers: [
        provideRouter([]),
        provideHttpClient(),
        provideHttpClientTesting(),
      ],
    });

    service = TestBed.inject(AuthService);
    httpMock = TestBed.inject(HttpTestingController);
    router = TestBed.inject(Router);
  });

  afterEach(() => {
    httpMock.verify();
    localStorage.clear();
  });

  it('stores token and user information after login', () => {
    service.login({ username: 'reader', password: 'secret123' }).subscribe();

    const req = httpMock.expectOne(`${environment.apiUrl}/login`);
    expect(req.request.method).toBe('POST');
    req.flush({
      message: 'Login successful',
      token:
        'eyJhbGciOiJIUzI1NiJ9.eyJ1c2VySWQiOjcsInJvbGUiOiJ1c2VyIn0.EcVro0Sm3jQ2h-9pk2NmM6rM4za9X6D8f9gC0xYIowQ',
      username: 'reader',
      role: 'user',
      userId: 7,
    });

    expect(service.isLoggedIn()).toBeTrue();
    expect(service.user()).toEqual({ id: 7, username: 'reader', role: 'user' });
    expect(localStorage.getItem('authToken')).toContain('.');
  });

  it('clears auth state and navigates to login on logout', () => {
    spyOn(router, 'navigate').and.resolveTo(true);
    localStorage.setItem('authToken', 'token');
    localStorage.setItem('user', JSON.stringify({ id: 3, username: 'admin', role: 'admin' }));

    service = TestBed.inject(AuthService);
    service.logout();

    expect(service.isLoggedIn()).toBeFalse();
    expect(localStorage.getItem('authToken')).toBeNull();
    expect(router.navigate).toHaveBeenCalledWith(['/login']);
  });
});
