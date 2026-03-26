import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { BookService } from '../../../core/services/book.service';
import { BorrowService } from '../../../core/services/borrow.service';
import { AuthService } from '../../../core/services/auth.service';
import { FavoritesService } from '../../../core/services/favorites.service';
import { RecentBooksService } from '../../../core/services/recent-books.service';
import { BookNotesService } from '../../../core/services/book-notes.service';
import { Book } from '../../../core/models/book.model';

@Component({
  selector: 'app-book-detail',
  standalone: true,
  imports: [RouterLink, FormsModule],
  templateUrl: './book-detail.component.html',
  styleUrl: './book-detail.component.css',
})
export class BookDetailComponent implements OnInit {
  book: Book | null = null;
  loading = true;
  actionLoading = false;
  message = '';
  isAdmin = false;
  noteText = '';

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private bookService: BookService,
    private borrowService: BorrowService,
    public auth: AuthService,
    public favorites: FavoritesService,
    private recent: RecentBooksService,
    private notes: BookNotesService,
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
      next: (b) => {
        this.book = b;
        this.recent.recordView({ id: b.id, title: b.title, author: b.author });
        this.noteText = this.notes.getNote(b.id);
      },
      error: () => {
        this.book = null;
        this.router.navigate(['/main/books']);
      },
      complete: () => (this.loading = false),
    });
  }

  saveNote(): void {
    if (this.book) {
      this.notes.setNote(this.book.id, this.noteText);
    }
  }

  toggleFavorite(): void {
    if (this.book) {
      this.favorites.toggle(this.book.id);
    }
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
