import { Routes } from '@angular/router';
import { authGuard } from './core/guards/auth.guard';
import { adminGuard } from './core/guards/admin.guard';

export const routes: Routes = [
  { path: '', redirectTo: 'login', pathMatch: 'full' },
  { path: 'login', loadComponent: () => import('./features/auth/login/login.component').then(m => m.LoginComponent) },
  { path: 'register', loadComponent: () => import('./features/auth/register/register.component').then(m => m.RegisterComponent) },
  {
    path: 'main',
    loadComponent: () => import('./layout/main-layout/main-layout.component').then(m => m.MainLayoutComponent),
    canActivate: [authGuard],
    children: [
      { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
      { path: 'dashboard', loadComponent: () => import('./features/dashboard/dashboard.component').then(m => m.DashboardComponent) },
      { path: 'books', loadComponent: () => import('./features/books/book-list/book-list.component').then(m => m.BookListComponent) },
      { path: 'books/:id', loadComponent: () => import('./features/books/book-detail/book-detail.component').then(m => m.BookDetailComponent) },
      { path: 'my-borrowings', loadComponent: () => import('./features/borrowings/my-borrowings/my-borrowings.component').then(m => m.MyBorrowingsComponent) },
      {
        path: 'admin/books',
        loadComponent: () => import('./features/admin/admin-books/admin-books.component').then(m => m.AdminBooksComponent),
        canActivate: [adminGuard],
      },
      {
        path: 'admin/books/new',
        loadComponent: () => import('./features/admin/book-form/book-form.component').then(m => m.BookFormComponent),
        canActivate: [adminGuard],
      },
      {
        path: 'admin/books/edit/:id',
        loadComponent: () => import('./features/admin/book-form/book-form.component').then(m => m.BookFormComponent),
        canActivate: [adminGuard],
      },
    ],
  },
  { path: '**', redirectTo: 'login' },
];
