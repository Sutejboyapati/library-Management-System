import { TestBed } from '@angular/core/testing';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { provideHttpClient } from '@angular/common/http';
import { BookService } from './book.service';

describe('BookService', () => {
  let service: BookService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [BookService, provideHttpClient(), provideHttpClientTesting()],
    });
    service = TestBed.inject(BookService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should call books endpoint with title search param', () => {
    service.getBooks('clean').subscribe();

    const req = httpMock.expectOne((r) => r.url.endsWith('/books') && r.params.get('title') === 'clean');
    expect(req.request.method).toBe('GET');
    req.flush([]);
  });

  it('should call book details endpoint by id', () => {
    service.getBook(7).subscribe();

    const req = httpMock.expectOne((r) => r.url.endsWith('/books/7'));
    expect(req.request.method).toBe('GET');
    req.flush({ id: 7 });
  });
});
