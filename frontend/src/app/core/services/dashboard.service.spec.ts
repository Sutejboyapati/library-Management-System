import { TestBed } from '@angular/core/testing';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting, HttpTestingController } from '@angular/common/http/testing';
import { DashboardService } from './dashboard.service';
import { environment } from '../../../environments/environment';

describe('DashboardService', () => {
  let service: DashboardService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [provideHttpClient(), provideHttpClientTesting()],
    });

    service = TestBed.inject(DashboardService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('requests the dashboard summary endpoint', () => {
    service.getSummary().subscribe((summary) => {
      expect(summary.totalBooks).toBe(12);
      expect(summary.availableBooks).toBe(20);
      expect(summary.activeBorrowings).toBe(4);
    });

    const req = httpMock.expectOne(`${environment.apiUrl}/dashboard/summary`);
    expect(req.request.method).toBe('GET');
    req.flush({
      totalBooks: 12,
      availableBooks: 20,
      activeBorrowings: 4,
    });
  });
});
