import { Injectable, signal } from '@angular/core';

const STORAGE_KEY = 'library_book_notes_v1';

@Injectable({ providedIn: 'root' })
export class BookNotesService {
  private readonly _notes = signal<Record<string, string>>({});
  readonly notes = this._notes.asReadonly();

  constructor() {
    this.hydrate();
  }

  private hydrate(): void {
    try {
      const raw = localStorage.getItem(STORAGE_KEY);
      this._notes.set(raw ? (JSON.parse(raw) as Record<string, string>) : {});
    } catch {
      this._notes.set({});
    }
  }

  private persist(): void {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(this._notes()));
  }

  getNote(bookId: number): string {
    return this._notes()[String(bookId)] ?? '';
  }

  setNote(bookId: number, text: string): void {
    const next = { ...this._notes() };
    const key = String(bookId);
    if (!text.trim()) {
      delete next[key];
    } else {
      next[key] = text.trim();
    }
    this._notes.set(next);
    this.persist();
  }
}
