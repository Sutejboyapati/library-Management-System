import { Component, OnInit } from '@angular/core';
import { DatePipe } from '@angular/common';
import { BorrowService } from '../../../core/services/borrow.service';
import { BorrowingRecord } from '../../../core/models/borrowing.model';

@Component({
  selector: 'app-my-borrowings',
  standalone: true,
  imports: [DatePipe],
  templateUrl: './my-borrowings.component.html',
  styleUrl: './my-borrowings.component.css',
})
export class MyBorrowingsComponent implements OnInit {
  records: BorrowingRecord[] = [];
  loading = true;
  actionLoading: number | null = null;
  message = '';

  constructor(private borrowService: BorrowService) {}

  ngOnInit(): void {
    this.load();
  }

  load(): void {
    this.loading = true;
    this.borrowService.getMyBorrowings().subscribe({
      next: (list) => {
        this.records = Array.isArray(list) ? list : [];
      },
      error: () => {
        this.records = [];
      },
      complete: () => (this.loading = false),
    });
  }

  returnBook(record: BorrowingRecord): void {
    const bookId = record.book_id;
    this.message = '';
    this.actionLoading = bookId;
    this.borrowService.returnBook(bookId).subscribe({
      next: () => {
        this.message = 'Book returned successfully.';
        this.load();
      },
      error: (err) => {
        this.message = err?.error ?? err?.message ?? 'Failed to return.';
      },
      complete: () => (this.actionLoading = null),
    });
  }

  activeRecords(): BorrowingRecord[] {
    return this.records.filter((r) => r.status === 'Borrowing' || !r.returned_at);
  }
}
