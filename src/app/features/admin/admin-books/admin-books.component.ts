import { Component, OnInit } from '@angular/core';
import { RouterLink } from '@angular/router';
import { BookService } from '../../../core/services/book.service';
import { Book } from '../../../core/models/book.model';

@Component({
  selector: 'app-admin-books',
  standalone: true,
  imports: [RouterLink],
  templateUrl: './admin-books.component.html',
  styleUrl: './admin-books.component.css',
})
export class AdminBooksComponent implements OnInit {
  books: Book[] = [];
  loading = false;
  deleteLoading: number | null = null;
  message = '';

  constructor(private bookService: BookService) {}

  ngOnInit(): void {
    this.load();
  }

  load(): void {
    this.loading = true;
    this.bookService.getBooks().subscribe({
      next: (list) => {
        this.books = Array.isArray(list) ? list : [];
      },
      error: () => {
        this.books = [];
      },
      complete: () => (this.loading = false),
    });
  }

  deleteBook(book: Book): void {
    if (!confirm(`Delete "${book.title}"?`)) return;
    this.message = '';
    this.deleteLoading = book.id;
    this.bookService.deleteBook(book.id).subscribe({
      next: () => {
        this.message = 'Book deleted.';
        this.load();
      },
      error: (err) => {
        this.message = err?.error?.message ?? err?.message ?? 'Delete failed.';
      },
      complete: () => (this.deleteLoading = null),
    });
  }
}
