import { TestBed } from '@angular/core/testing';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting, HttpTestingController } from '@angular/common/http/testing';
import { BookService } from './book.service';
import { environment } from '../../../environments/environment';

describe('BookService', () => {
  let service: BookService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [provideHttpClient(), provideHttpClientTesting()],
    });

    service = TestBed.inject(BookService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('sends the search term through the q query parameter', () => {
    service.getBooks('clean code').subscribe();

    const req = httpMock.expectOne(
      (request) => request.url === `${environment.apiUrl}/books` && request.params.get('q') === 'clean code',
    );
    expect(req.request.method).toBe('GET');
    req.flush([]);
  });

  it('loads a single book by id', () => {
    service.getBook(12).subscribe();

    const req = httpMock.expectOne(`${environment.apiUrl}/books/12`);
    expect(req.request.method).toBe('GET');
    req.flush({ id: 12, title: 'Clean Code', author: 'Robert C. Martin', available_copies: 3 });
  });
});
