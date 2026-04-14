import { TestBed } from '@angular/core/testing';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting, HttpTestingController } from '@angular/common/http/testing';
import { environment } from '../../../environments/environment';
import { ReviewService } from './review.service';

describe('ReviewService', () => {
  let service: ReviewService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [provideHttpClient(), provideHttpClientTesting()],
    });

    service = TestBed.inject(ReviewService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('loads reviews for a book', () => {
    service.getReviews(5).subscribe();

    const req = httpMock.expectOne(`${environment.apiUrl}/books/5/reviews`);
    expect(req.request.method).toBe('GET');
    req.flush([]);
  });

  it('posts a review for a book', () => {
    service.saveReview(5, { rating: 4, comment: 'Helpful and practical.' }).subscribe();

    const req = httpMock.expectOne(`${environment.apiUrl}/books/5/reviews`);
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({ rating: 4, comment: 'Helpful and practical.' });
    req.flush({
      id: 1,
      bookId: 5,
      userId: 2,
      username: 'reader',
      rating: 4,
      comment: 'Helpful and practical.',
      createdAt: '2026-04-13T10:00:00Z',
      updatedAt: '2026-04-13T10:00:00Z',
    });
  });
});
