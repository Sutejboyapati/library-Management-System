import { Component, OnInit } from '@angular/core';
import { RouterLink } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { BookService } from '../../../core/services/book.service';
import { Book } from '../../../core/models/book.model';

@Component({
  selector: 'app-book-list',
  standalone: true,
  imports: [RouterLink, FormsModule],
  templateUrl: './book-list.component.html',
  styleUrl: './book-list.component.css',
})
export class BookListComponent implements OnInit {
  books: Book[] = [];
  loading = false;
  search = '';

  constructor(private bookService: BookService) {}

  ngOnInit(): void {
    this.loadBooks();
  }

  loadError = '';

  loadBooks(): void {
    this.loading = true;
    this.loadError = '';
    this.bookService.getBooks(this.search || undefined).subscribe({
      next: (list) => {
        this.books = Array.isArray(list) ? list : [];
      },
      error: (err) => {
        this.books = [];
        this.loadError = err?.status === 0
          ? 'Cannot reach backend. Make sure it is running on port 3000.'
          : 'Failed to load books.';
      },
      complete: () => (this.loading = false),
    });
  }

  onSearch(): void {
    this.loadBooks();
  }
}
