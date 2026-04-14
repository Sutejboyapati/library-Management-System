import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';
import { Review, ReviewRequest } from '../models/review.model';

@Injectable({ providedIn: 'root' })
export class ReviewService {
  constructor(private http: HttpClient) {}

  getReviews(bookId: number): Observable<Review[]> {
    return this.http.get<Review[]>(`${environment.apiUrl}/books/${bookId}/reviews`);
  }

  saveReview(bookId: number, payload: ReviewRequest): Observable<Review> {
    return this.http.post<Review>(`${environment.apiUrl}/books/${bookId}/reviews`, payload);
  }
}
