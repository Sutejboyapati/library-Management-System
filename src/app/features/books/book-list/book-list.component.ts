import { Component, OnInit } from '@angular/core';
import { RouterLink } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { BookService } from '../../../core/services/book.service';
import { FavoritesService } from '../../../core/services/favorites.service';
import { Book } from '../../../core/models/book.model';

export type BookSort = 'title' | 'titleDesc' | 'author' | 'copies';

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
  loadError = '';

  sortBy: BookSort = 'title';
  genreFilter = '';
  inStockOnly = false;
  favoritesOnly = false;
  viewMode: 'grid' | 'list' = 'grid';

  constructor(
    private bookService: BookService,
    public favorites: FavoritesService,
  ) {}

  ngOnInit(): void {
    this.loadBooks();
  }

  loadBooks(): void {
    this.loading = true;
    this.loadError = '';
    this.bookService.getBooks(this.search || undefined).subscribe({
      next: (list) => {
        this.books = Array.isArray(list) ? list : [];
      },
      error: (err) => {
        this.books = [];
        this.loadError =
          err?.status === 0
            ? 'Cannot reach backend. Make sure it is running on port 3000.'
            : 'Failed to load books.';
      },
      complete: () => (this.loading = false),
    });
  }

  onSearch(): void {
    this.loadBooks();
  }

  get genres(): string[] {
    const g = new Set<string>();
    for (const b of this.books) {
      const v = (b.genre ?? '').trim();
      if (v) {
        g.add(v);
      }
    }
    return [...g].sort((a, b) => a.localeCompare(b, undefined, { sensitivity: 'base' }));
  }

  get filteredBooks(): Book[] {
    let list = [...this.books];
    if (this.inStockOnly) {
      list = list.filter((b) => b.available_copies > 0);
    }
    if (this.genreFilter) {
      list = list.filter((b) => (b.genre ?? '').trim() === this.genreFilter);
    }
    if (this.favoritesOnly) {
      const ids = new Set(this.favorites.ids());
      list = list.filter((b) => ids.has(b.id));
    }
    const cmp = (a: string, b: string) => a.localeCompare(b, undefined, { sensitivity: 'base' });
    switch (this.sortBy) {
      case 'title':
        list.sort((a, b) => cmp(a.title, b.title));
        break;
      case 'titleDesc':
        list.sort((a, b) => cmp(b.title, a.title));
        break;
      case 'author':
        list.sort((a, b) => cmp(a.author, b.author));
        break;
      case 'copies':
        list.sort((a, b) => b.available_copies - a.available_copies);
        break;
    }
    return list;
  }

  toggleFavorite(event: Event, id: number): void {
    event.preventDefault();
    event.stopPropagation();
    this.favorites.toggle(id);
  }
}
