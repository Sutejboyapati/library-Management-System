import { Component, OnInit } from '@angular/core';
import { DatePipe } from '@angular/common';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { BookService } from '../../../core/services/book.service';
import { BorrowService } from '../../../core/services/borrow.service';
import { AuthService } from '../../../core/services/auth.service';
import { Book } from '../../../core/models/book.model';
import { Review } from '../../../core/models/review.model';
import { ReviewService } from '../../../core/services/review.service';

@Component({
  selector: 'app-il',
  standalone: true,
  imports: [RouterLink, FormsModule, DatePipe],
  templateUrl: './il.component.html',
  styleUrl: './il.component.css',
})
export class BookDetailComponent implements OnInit {
  book: Book | null = null;
  reviews: Review[] = [];
  loading = true;
  reviewsLoading = false;
  actionLoading = false;
  reviewLoading = false;
  message = '';
  reviewMessage = '';
  reviewError = '';
  isAdmin = false;
  reviewRating = 5;
  reviewComment = '';

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private bookService: BookService,
    private borrowService: BorrowService,
    private reviewService: ReviewService,
    public auth: AuthService,
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
        this.loadReviews(numId);
      },
      error: () => {
        this.book = null;
        this.router.navigate(['/main/books']);
      },
      complete: () => (this.loading = false),
    });
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

  loadReviews(bookId: number): void {
    this.reviewsLoading = true;
    this.reviewError = '';
    this.reviewService.getReviews(bookId).subscribe({
      next: (reviews) => {
        this.reviews = Array.isArray(reviews) ? reviews : [];
        this.prefillReviewForm();
      },
      error: () => {
        this.reviews = [];
        this.reviewError = 'Failed to load reviews.';
      },
      complete: () => {
        this.reviewsLoading = false;
      },
    });
  }

  submitReview(): void {
    if (!this.book || !this.canReview || this.reviewLoading) return;

    this.reviewMessage = '';
    this.reviewError = '';

    if (!this.reviewComment.trim()) {
      this.reviewError = 'Please write a short review comment.';
      return;
    }

    this.reviewLoading = true;
    this.reviewService.saveReview(this.book.id, {
      rating: this.reviewRating,
      comment: this.reviewComment.trim(),
    }).subscribe({
      next: () => {
        this.reviewMessage = 'Review saved successfully.';
        this.loadReviews(this.book!.id);
      },
      error: (err) => {
        this.reviewError = err?.error?.message ?? err?.message ?? 'Failed to save review.';
      },
      complete: () => {
        this.reviewLoading = false;
      },
    });
  }

  prefillReviewForm(): void {
    const currentUser = this.auth.user();
    if (!currentUser) return;

    const existingReview = this.reviews.find((review) => review.userId === Number(currentUser.id));
    if (!existingReview) return;

    this.reviewRating = existingReview.rating;
    this.reviewComment = existingReview.comment;
  }

  get available(): number {
    return this.book?.available_copies ?? 0;
  }

  get averageRating(): string {
    if (!this.reviews.length) return 'No ratings yet';
    const total = this.reviews.reduce((sum, review) => sum + review.rating, 0);
    return (total / this.reviews.length).toFixed(1);
  }

  get canReview(): boolean {
    return this.auth.isLoggedIn() && !this.isAdmin;
  }
}
