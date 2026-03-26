import { Injectable, signal } from '@angular/core';

const STORAGE_KEY = 'library_favorites_v1';

@Injectable({ providedIn: 'root' })
export class FavoritesService {
  private readonly _ids = signal<number[]>([]);
  readonly ids = this._ids.asReadonly();

  constructor() {
    this.hydrate();
  }

  private hydrate(): void {
    try {
      const raw = localStorage.getItem(STORAGE_KEY);
      this._ids.set(raw ? (JSON.parse(raw) as number[]) : []);
    } catch {
      this._ids.set([]);
    }
  }

  private persist(): void {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(this._ids()));
  }

  isFavorite(id: number): boolean {
    return this._ids().includes(id);
  }

  toggle(id: number): void {
    const cur = this._ids();
    const next = cur.includes(id) ? cur.filter((x) => x !== id) : [...cur, id];
    this._ids.set(next);
    this.persist();
  }

  count(): number {
    return this._ids().length;
  }
}
