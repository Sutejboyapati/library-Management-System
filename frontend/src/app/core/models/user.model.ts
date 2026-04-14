export interface User {
  id: number | string;
  username: string;
  role: 'user' | 'admin';
}
