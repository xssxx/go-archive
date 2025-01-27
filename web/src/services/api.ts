import { Book } from "../types/book";

const API_URL = "http://localhost:3000";

export async function fetchBooks(): Promise<Book[]> {
  try {
    const response = await fetch(`${API_URL}/books`);
    if (!response.ok) throw new Error("Failed to fetch books");
    const books: Book[] = await response.json();

    return books;
  } catch (error) {
    console.error(error);

    return [];
  }
}
