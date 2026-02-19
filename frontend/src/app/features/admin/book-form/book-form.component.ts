import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { BookService } from '../../../core/services/book.service';
import { Book } from '../../../core/models/book.model';

@Component({
  selector: 'app-book-form',
  standalone: true,
  imports: [FormsModule, RouterLink],
  templateUrl: './book-form.component.html',
  styleUrl: './book-form.component.css',
})
export class BookFormComponent implements OnInit {
  isEdit = false;
  id: number | null = null;
  title = '';
  author = '';
  isbn = '';
  genre = '';
  language = 'English';
  shelf_number = '';
  available_copies = 1;
  loading = false;
  submitLoading = false;
  error = '';

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private bookService: BookService,
  ) {}

  ngOnInit(): void {
    const idParam = this.route.snapshot.paramMap.get('id');
    if (idParam) {
      this.isEdit = true;
      this.id = Number(idParam);
      if (!Number.isNaN(this.id)) this.loadBook();
    }
  }

  loadBook(): void {
    if (this.id == null) return;
    this.loading = true;
    this.bookService.getBook(this.id).subscribe({
      next: (b) => {
        this.title = b.title ?? '';
        this.author = b.author ?? '';
        this.isbn = b.isbn ?? '';
        this.genre = b.genre ?? '';
        this.language = b.language ?? 'English';
        this.shelf_number = b.shelf_number ?? '';
        this.available_copies = b.available_copies ?? 1;
      },
      error: () => this.router.navigate(['/main/admin/books']),
      complete: () => (this.loading = false),
    });
  }

  onSubmit(): void {
    this.error = '';
    if (!this.title.trim()) {
      this.error = 'Title is required.';
      return;
    }
    if (!this.author.trim()) {
      this.error = 'Author is required.';
      return;
    }
    const payload = {
      title: this.title.trim(),
      author: this.author.trim(),
      isbn: this.isbn.trim() || undefined,
      genre: this.genre.trim() || undefined,
      language: this.language.trim() || undefined,
      shelf_number: this.shelf_number.trim() || undefined,
      available_copies: this.available_copies,
    };
    this.submitLoading = true;
    if (this.isEdit && this.id != null) {
      this.bookService.updateBook(this.id, payload).subscribe({
        next: () => this.router.navigate(['/main/admin/books']),
        error: (err) => {
          this.error = err?.error?.message ?? err?.message ?? 'Update failed.';
          this.submitLoading = false;
        },
        complete: () => (this.submitLoading = false),
      });
    } else {
      this.bookService.addBook(payload).subscribe({
        next: () => this.router.navigate(['/main/admin/books']),
        error: (err) => {
          this.error = err?.error?.message ?? err?.message ?? 'Add failed.';
          this.submitLoading = false;
        },
        complete: () => (this.submitLoading = false),
      });
    }
  }
}
