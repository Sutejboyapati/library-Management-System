import { Component, OnInit } from '@angular/core';
import { RouterLink } from '@angular/router';
import { BookService } from '../../core/services/book.service';
import { FavoritesService } from '../../core/services/favorites.service';
import { Book } from '../../core/models/book.model';

@Component({
  selector: 'app-favorites',
  standalone: true,
  imports: [RouterLink],
  templateUrl: './favorites.component.html',
  styleUrl: './favorites.component.css',
})
export class FavoritesComponent implements OnInit {
  books: Book[] = [];
  loading = true;

  constructor(
    private bookService: BookService,
    public favorites: FavoritesService,
  ) {}

  ngOnInit(): void {
    this.load();
  }

  load(): void {
    const ids = this.favorites.ids();
    if (ids.length === 0) {
      this.books = [];
      this.loading = false;
      return;
    }
    this.loading = true;
    this.bookService.getBooks().subscribe({
      next: (list) => {
        const all = Array.isArray(list) ? list : [];
        const set = new Set(ids);
        this.books = all.filter((b) => set.has(b.id));
      },
      error: () => {
        this.books = [];
      },
      complete: () => (this.loading = false),
    });
  }

  toggleFavorite(event: Event, id: number): void {
    event.preventDefault();
    event.stopPropagation();
    this.favorites.toggle(id);
    this.load();
  }
}
