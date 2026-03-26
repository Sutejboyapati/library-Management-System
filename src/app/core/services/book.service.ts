import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';
import { Book } from '../models/book.model';

@Injectable({ providedIn: 'root' })
export class BookService {
  private base = `${environment.apiUrl}/books`;

  constructor(private http: HttpClient) {}

  getBooks(search?: string): Observable<Book[]> {
    let params = new HttpParams();
    if (search?.trim()) params = params.set('title', search.trim());
    return this.http.get<Book[]>(this.base, { params });
  }

  getBook(id: number): Observable<Book> {
    return this.http.get<Book>(`${this.base}/${id}`);
  }

  addBook(book: Partial<Book>): Observable<{ message: string }> {
    return this.http.post<{ message: string }>(this.base, book);
  }

  updateBook(id: number, book: Partial<Book>): Observable<{ message: string }> {
    return this.http.put<{ message: string }>(`${this.base}/${id}`, book);
  }

  deleteBook(id: number): Observable<{ message: string }> {
    return this.http.delete<{ message: string }>(`${this.base}/${id}`);
  }
}
