import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideRouter } from '@angular/router';
import { of } from 'rxjs';

import { LoginComponent } from './login.component';
import { AuthService } from '../../../core/services/auth.service';

describe('LoginComponent', () => {
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;
  let authSpy: jasmine.SpyObj<AuthService>;

  beforeEach(async () => {
    authSpy = jasmine.createSpyObj<AuthService>('AuthService', ['login', 'isAdmin'], {
      isLoggedIn: () => false,
      user: () => null,
    } as Partial<AuthService>);

    authSpy.login.and.returnValue(of({ message: 'ok', token: 'token' }));
    authSpy.isAdmin.and.returnValue(false);

    await TestBed.configureTestingModule({
      imports: [LoginComponent],
      providers: [
        provideRouter([]),
        { provide: AuthService, useValue: authSpy },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should show validation error when username/password are empty', () => {
    component.username = '';
    component.password = '';

    component.onSubmit();

    expect(component.error).toContain('Please enter username and password');
    expect(authSpy.login).not.toHaveBeenCalled();
  });

  it('should call auth login on valid submit', () => {
    component.username = 'admin';
    component.password = 'admin123';

    component.onSubmit();

    expect(authSpy.login).toHaveBeenCalledWith({ username: 'admin', password: 'admin123' });
  });
});
