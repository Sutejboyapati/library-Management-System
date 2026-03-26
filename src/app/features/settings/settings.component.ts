import { Component } from '@angular/core';
import { BorrowService } from '../../core/services/borrow.service';
import { AccessibilityService, FontScale } from '../../core/services/accessibility.service';
import { BorrowingRecord } from '../../core/models/borrowing.model';

@Component({
  selector: 'app-settings',
  standalone: true,
  imports: [],
  templateUrl: './settings.component.html',
  styleUrl: './settings.component.css',
})
export class SettingsComponent {
  exportMessage = '';
  exportLoading = false;

  constructor(
    public a11y: AccessibilityService,
    private borrowService: BorrowService,
  ) {}

  setFont(v: FontScale): void {
    this.a11y.setFontScale(v);
  }

  setContrast(v: boolean): void {
    this.a11y.setHighContrast(v);
  }

  setMotion(v: boolean): void {
    this.a11y.setReduceMotion(v);
  }

  exportBorrowingHistory(): void {
    this.exportMessage = '';
    this.exportLoading = true;
    this.borrowService.getMyBorrowings().subscribe({
      next: (list) => {
        const rows = Array.isArray(list) ? list : [];
        const csv = this.toCsv(rows);
        const blob = new Blob([csv], { type: 'text/csv;charset=utf-8' });
        const url = URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        a.download = `library-borrowing-history-${new Date().toISOString().slice(0, 10)}.csv`;
        a.click();
        URL.revokeObjectURL(url);
        this.exportMessage = 'Download started.';
      },
      error: () => {
        this.exportMessage = 'Could not load borrowings. Try again when the backend is running.';
      },
      complete: () => (this.exportLoading = false),
    });
  }

  private toCsv(rows: BorrowingRecord[]): string {
    const header = ['Book ID', 'Title', 'Author', 'ISBN', 'Borrowed', 'Due', 'Returned', 'Status'];
    const lines = [header.join(',')];
    for (const r of rows) {
      const cells = [
        String(r.book_id ?? ''),
        this.escapeCsv(r.title ?? ''),
        this.escapeCsv(r.author ?? ''),
        this.escapeCsv(r.isbn ?? ''),
        r.borrowed_at ? new Date(r.borrowed_at).toISOString() : '',
        r.due_date ? new Date(r.due_date).toISOString() : '',
        r.returned_at ? new Date(r.returned_at).toISOString() : '',
        this.escapeCsv(r.status ?? ''),
      ];
      lines.push(cells.join(','));
    }
    return lines.join('\r\n');
  }

  private escapeCsv(s: string): string {
    if (s.includes(',') || s.includes('"') || s.includes('\n')) {
      return `"${s.replace(/"/g, '""')}"`;
    }
    return s;
  }
}
