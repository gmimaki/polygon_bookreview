import React, { useEffect, useMemo, useState } from "react";
import { getBooks } from "../util/interact";

const BookList = () => {
  	const [books, setBooks] = useState([])

	useEffect(() => {
		async function fetchBooks() {
		  const books = await getBooks();
		  setBooks(books)
		  console.log(books)
		}
		fetchBooks();
	}, []);

	const booksElm = useMemo(() => {
		return books.map(b => {
		  return (
		    <>
		      <li>{b.name} ({b.author})</li>
		    </>
		  )
		})
	}, [books])

	return (
		<>
			<ul>
				{booksElm}
			</ul>
		</>
	)
}

export default BookList