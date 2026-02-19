export interface Book {
  id: number;
  title: string;
  author: string;
  isbn?: string;
  genre?: string;
  language?: string;
  shelf_number?: string;
  available_copies: number;
}
