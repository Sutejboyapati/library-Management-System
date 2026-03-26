import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';
import { AuthService } from './auth.service';
import { BorrowingRecord } from '../models/borrowing.model';

@Injectable({ providedIn: 'root' })
export class BorrowService {
  private api = environment.apiUrl;

  constructor(
    private http: HttpClient,
    private auth: AuthService,
  ) {}

  borrow(bookId: number): Observable<{ message: string }> {
    const userId = this.auth.getUserId();
    return this.http.post<{ message: string }>(`${this.api}/borrow`, {
      userId: userId ?? 0,
      bookId,
    });
  }

  returnBook(bookId: number): Observable<{ message: string }> {
    const userId = this.auth.getUserId();
    return this.http.post<{ message: string }>(`${this.api}/borrow/return`, {
      userId: userId ?? 0,
      bookId,
    });
  }

  getMyBorrowings(): Observable<BorrowingRecord[]> {
    const userId = this.auth.getUserId();
    if (userId == null) return new Observable((o) => o.next([]));
    return this.http.get<BorrowingRecord[]>(`${this.api}/users/${userId}/borrowings`);
  }
}
