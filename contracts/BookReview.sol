// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

// import "hardhat/console.sol";

contract BookReview {
    event NewBook(uint bookId, string name, string author);

    struct Book {
        string name;
        string author;
    }

    Book[] public books;

    function registerBook(string memory _name, string memory _author) public {
        books.push(Book(_name, _author));
        uint bookId = books.length - 1;
        emit NewBook(bookId, _name, _author);
    }

    function findBook() public view returns (Book[] memory) {
        return books;
    }

    event NewReview(uint reviewId, uint bookId, string contenxt);

    struct Review {
        uint bookId;
        string content;
    }

    Review[] public reviews;

    function registerReview(uint _bookId, string memory _content) public {
        reviews.push(Review(_bookId, _content));
        uint reviewId = reviews.length - 1;
        emit NewReview(reviewId, _bookId, _content);
    }

    function findReview(uint _reviewId) public view returns (Review memory) {
        return reviews[_reviewId];
    }
}