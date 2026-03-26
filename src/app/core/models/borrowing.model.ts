export interface BorrowingRecord {
  id?: number;
  user_id?: number;
  book_id: number;
  title?: string;
  author?: string;
  isbn?: string;
  borrowed_at: string;
  due_date: string;
  returned_at?: string | null;
  status?: string;
}
