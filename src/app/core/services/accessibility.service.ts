import { Injectable, signal } from '@angular/core';

export type FontScale = 'normal' | 'large' | 'xlarge';

@Injectable({ providedIn: 'root' })
export class AccessibilityService {
  private readonly root = document.documentElement;

  readonly fontScale = signal<FontScale>('normal');
  readonly highContrast = signal(false);
  readonly reduceMotion = signal(false);

  constructor() {
    this.loadFromStorage();
    this.apply();
  }

  private loadFromStorage(): void {
    try {
      const f = localStorage.getItem('a11y_font') as FontScale | null;
      if (f === 'large' || f === 'xlarge' || f === 'normal') {
        this.fontScale.set(f);
      }
      this.highContrast.set(localStorage.getItem('a11y_contrast') === '1');
      this.reduceMotion.set(localStorage.getItem('a11y_motion') === '1');
    } catch {
      /* ignore */
    }
  }

  setFontScale(v: FontScale): void {
    this.fontScale.set(v);
    localStorage.setItem('a11y_font', v);
    this.apply();
  }

  setHighContrast(v: boolean): void {
    this.highContrast.set(v);
    localStorage.setItem('a11y_contrast', v ? '1' : '0');
    this.apply();
  }

  setReduceMotion(v: boolean): void {
    this.reduceMotion.set(v);
    localStorage.setItem('a11y_motion', v ? '1' : '0');
    this.apply();
  }

  private apply(): void {
    this.root.classList.remove('a11y-font-large', 'a11y-font-xlarge', 'a11y-high-contrast', 'a11y-reduce-motion');
    const f = this.fontScale();
    if (f === 'large') {
      this.root.classList.add('a11y-font-large');
    } else if (f === 'xlarge') {
      this.root.classList.add('a11y-font-xlarge');
    }
    if (this.highContrast()) {
      this.root.classList.add('a11y-high-contrast');
    }
    if (this.reduceMotion()) {
      this.root.classList.add('a11y-reduce-motion');
    }
  }
}
