import { useEffect, useState } from "react";
import { fetchBooks } from "./services/api";
import { type Book } from "./types/book";

function App() {
  const [books, setBooks] = useState<Book[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>("");

  useEffect(() => {
    const getBooks = async () => {
      try {
        const fetchedBooks = await fetchBooks();
        setBooks(fetchedBooks);
      } catch (error) {
        setError("Failed to fetch books");
      } finally {
        setLoading(false);
      }
    };

    getBooks();
  }, []);

  return (
    <div>
      <h1>Book List</h1>
      {loading ? (
        <p>Loading...</p>
      ) : error ? (
        <p>{error}</p>
      ) : (
        <ul>
          {books.map((book) => (
            <li key={book.id}>
              <strong>{book.title}</strong> by {book.author} (
              {book.publishedDate})
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default App;
