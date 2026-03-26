import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { BookService } from '../../../core/services/book.service';
import { BorrowService } from '../../../core/services/borrow.service';
import { AuthService } from '../../../core/services/auth.service';
import { Book } from '../../../core/models/book.model';

@Component({
  selector: 'app-book-detail',
  standalone: true,
  imports: [RouterLink],
  templateUrl: './book-detail.component.html',
  styleUrl: './book-detail.component.css',
})
export class BookDetailComponent implements OnInit {
  book: Book | null = null;
  loading = true;
  actionLoading = false;
  message = '';
  isAdmin = false;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private bookService: BookService,
    private borrowService: BorrowService,
    public auth: AuthService,
  ) {
    this.isAdmin = this.auth.isAdmin();
  }

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    if (!id) {
      this.router.navigate(['/main/books']);
      return;
    }
    const numId = Number(id);
    if (Number.isNaN(numId)) {
      this.router.navigate(['/main/books']);
      return;
    }
    this.loading = true;
    this.bookService.getBook(numId).subscribe({
      next: (b) => (this.book = b),
      error: () => {
        this.book = null;
        this.router.navigate(['/main/books']);
      },
      complete: () => (this.loading = false),
    });
  }

  borrow(): void {
    if (!this.book || this.actionLoading) return;
    this.message = '';
    this.actionLoading = true;
    this.borrowService.borrow(this.book.id).subscribe({
      next: () => {
        this.message = 'Book borrowed successfully.';
        this.refreshBook();
      },
      error: (err) => {
        this.message = err?.error ?? err?.message ?? 'Failed to borrow.';
      },
      complete: () => (this.actionLoading = false),
    });
  }

  private refreshBook(): void {
    if (!this.book) return;
    this.bookService.getBook(this.book.id).subscribe({
      next: (b) => (this.book = b),
    });
  }

  get available(): number {
    return this.book?.available_copies ?? 0;
  }
}
