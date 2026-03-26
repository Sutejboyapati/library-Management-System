import { Injectable, signal } from '@angular/core';

const STORAGE_KEY = 'library_recent_v1';

export interface RecentBookEntry {
  id: number;
  title: string;
  author: string;
  visitedAt: number;
}

@Injectable({ providedIn: 'root' })
export class RecentBooksService {
  private readonly _recent = signal<RecentBookEntry[]>([]);
  readonly recent = this._recent.asReadonly();

  constructor() {
    this.hydrate();
  }

  private hydrate(): void {
    try {
      const raw = localStorage.getItem(STORAGE_KEY);
      this._recent.set(raw ? (JSON.parse(raw) as RecentBookEntry[]) : []);
    } catch {
      this._recent.set([]);
    }
  }

  recordView(book: { id: number; title: string; author: string }): void {
    const without = this._recent().filter((e) => e.id !== book.id);
    const entry: RecentBookEntry = {
      ...book,
      visitedAt: Date.now(),
    };
    const next = [entry, ...without].slice(0, 12);
    this._recent.set(next);
    localStorage.setItem(STORAGE_KEY, JSON.stringify(next));
  }
}
