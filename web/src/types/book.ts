export interface Book {
  id: number;
  title: string;
  author: string;
  publishedDate?: string;
  genre?: string;
  price: number;
  stock: number;
  createdAt?: string;
  updatedAt?: string;
}
