package gormsql

import (
	"fmt"
	"os"

	"github.com/anggaaryas/go-mockapi"
	"gorm.io/gorm"
)

type dataSource struct {
	db *gorm.DB
}

func getBaseURL() string {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	return baseURL
}

func getCoverURL(filename string) string {
	return fmt.Sprintf("%s%s%s", getBaseURL(), mockapi.GetMockapiStaticImagePath(), filename)
}

func Create(db *gorm.DB) mockapi.DataSource {
	return &dataSource{
		db: db,
	}
}

func getInitialBooks() []mockapi.Book {
	var books = []mockapi.Book{
		{ID: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Category: "Programming", Desc: "A comprehensive guide to Go programming", CoverURL: getCoverURL("go-programming-languange.jpg")},
		{ID: 2, Title: "Clean Code", Author: "Robert C. Martin", Category: "Programming", Desc: "A handbook of agile software craftsmanship", CoverURL: getCoverURL("clean-code.jpg")},
		{ID: 3, Title: "Design Patterns", Author: "Erich Gamma", Category: "Programming", Desc: "Elements of reusable object-oriented software", CoverURL: getCoverURL("design-pattern.jpg")},
		{ID: 4, Title: "The Pragmatic Programmer", Author: "Andrew Hunt", Category: "Programming", Desc: "Your journey to mastery", CoverURL: getCoverURL("the-pragmatic-programmer.jpg")},
		{ID: 5, Title: "Introduction to Algorithms", Author: "Thomas H. Cormen", Category: "Computer Science", Desc: "A comprehensive introduction to algorithms", CoverURL: getCoverURL("introduction-to-algorithms.jpg")},
		{ID: 6, Title: "Code Complete", Author: "Steve McConnell", Category: "Programming", Desc: "A practical handbook of software construction", CoverURL: getCoverURL("code-complete.jpg")},
		{ID: 7, Title: "Refactoring", Author: "Martin Fowler", Category: "Programming", Desc: "Improving the design of existing code", CoverURL: getCoverURL("refactoring.jpg")},
		{ID: 8, Title: "Head First Design Patterns", Author: "Eric Freeman", Category: "Programming", Desc: "A brain-friendly guide to design patterns", CoverURL: getCoverURL("head-first-design-pattern.jpg")},
		{ID: 9, Title: "The Mythical Man-Month", Author: "Frederick P. Brooks Jr.", Category: "Software Engineering", Desc: "Essays on software engineering", CoverURL: getCoverURL("the-mythical-man-month.jpg")},
		{ID: 10, Title: "Cracking the Coding Interview", Author: "Gayle Laakmann McDowell", Category: "Programming", Desc: "189 programming questions and solutions", CoverURL: getCoverURL("cracking-the-code-interview.jpg")},
		{ID: 11, Title: "You Don't Know JS", Author: "Kyle Simpson", Category: "Programming", Desc: "Deep dive into JavaScript", CoverURL: getCoverURL("you-dont-know-js.jpg")},
		{ID: 12, Title: "Eloquent JavaScript", Author: "Marijn Haverbeke", Category: "Programming", Desc: "A modern introduction to programming", CoverURL: getCoverURL("eloquent-javascript.jpg")},
		{ID: 13, Title: "JavaScript: The Good Parts", Author: "Douglas Crockford", Category: "Programming", Desc: "Unearthing the excellence in JavaScript", CoverURL: getCoverURL("javascript-the-good-parts.jpg")},
		{ID: 14, Title: "The Art of Computer Programming", Author: "Donald Knuth", Category: "Computer Science", Desc: "Fundamental algorithms", CoverURL: getCoverURL("the-art-of-computer-programming.jpg")},
		{ID: 15, Title: "Structure and Interpretation of Computer Programs", Author: "Harold Abelson", Category: "Computer Science", Desc: "Classic computer science text", CoverURL: getCoverURL("structure-and-interpretation-of-computer-programs.jpg")},
		{ID: 16, Title: "Python Crash Course", Author: "Eric Matthes", Category: "Programming", Desc: "A hands-on project-based introduction to programming", CoverURL: getCoverURL("python-crash-course.jpg")},
		{ID: 17, Title: "Learning Python", Author: "Mark Lutz", Category: "Programming", Desc: "Powerful object-oriented programming", CoverURL: getCoverURL("learning-python.jpg")},
		{ID: 18, Title: "Fluent Python", Author: "Luciano Ramalho", Category: "Programming", Desc: "Clear, concise, and effective programming", CoverURL: getCoverURL("fluent-python.jpg")},
		{ID: 19, Title: "Automate the Boring Stuff with Python", Author: "Al Sweigart", Category: "Programming", Desc: "Practical programming for total beginners", CoverURL: getCoverURL("automate-the-boring-stuff-with-python.jpg")},
		{ID: 20, Title: "Effective Java", Author: "Joshua Bloch", Category: "Programming", Desc: "Best practices for the Java platform", CoverURL: getCoverURL("effective-java.jpg")},
		{ID: 21, Title: "Java: The Complete Reference", Author: "Herbert Schildt", Category: "Programming", Desc: "Comprehensive guide to Java programming", CoverURL: getCoverURL("java-the-complete-reference.jpg")},
		{ID: 22, Title: "Head First Java", Author: "Kathy Sierra", Category: "Programming", Desc: "A brain-friendly guide to Java", CoverURL: getCoverURL("head-first-java.jpg")},
		{ID: 23, Title: "Thinking in Java", Author: "Bruce Eckel", Category: "Programming", Desc: "The definitive introduction to Java", CoverURL: getCoverURL("thinking-in-java.jpg")},
		{ID: 24, Title: "C Programming Language", Author: "Brian Kernighan", Category: "Programming", Desc: "The classic C programming guide", CoverURL: getCoverURL("c-programming-language.jpg")},
		{ID: 25, Title: "C++ Primer", Author: "Stanley Lippman", Category: "Programming", Desc: "Comprehensive introduction to C++", CoverURL: getCoverURL("Cpp-Primer.jpg")},
		{ID: 26, Title: "Effective Modern C++", Author: "Scott Meyers", Category: "Programming", Desc: "42 specific ways to improve your use of C++11 and C++14", CoverURL: getCoverURL("effective-modern-cpp.jpg")},
		{ID: 27, Title: "The C++ Programming Language", Author: "Bjarne Stroustrup", Category: "Programming", Desc: "The definitive guide by the creator of C++", CoverURL: getCoverURL("the-cpp-programming-language.jpg")},
		{ID: 28, Title: "Ruby on Rails Tutorial", Author: "Michael Hartl", Category: "Web Development", Desc: "Learn web development with Rails", CoverURL: getCoverURL("ruby-on-rails-tutorial.jpg")},
		{ID: 29, Title: "Programming Ruby", Author: "Dave Thomas", Category: "Programming", Desc: "The pragmatic programmers guide", CoverURL: getCoverURL("programming-ruby.jpg")},
		{ID: 30, Title: "Node.js Design Patterns", Author: "Mario Casciaro", Category: "Web Development", Desc: "Master best practices to build modular applications", CoverURL: getCoverURL("node-js-design-patterns.jpg")},
		{ID: 31, Title: "Learning React", Author: "Alex Banks", Category: "Web Development", Desc: "Modern patterns for developing React apps", CoverURL: getCoverURL("learning-react.jpg")},
		{ID: 32, Title: "React Up & Running", Author: "Stoyan Stefanov", Category: "Web Development", Desc: "Building web applications with React", CoverURL: getCoverURL("react-up-running.jpg")},
		{ID: 33, Title: "Vue.js in Action", Author: "Erik Hanchett", Category: "Web Development", Desc: "Building modern web applications with Vue", CoverURL: getCoverURL("vue-js-in-action.jpg")},
		{ID: 34, Title: "Angular in Action", Author: "Jeremy Wilken", Category: "Web Development", Desc: "Build dynamic web applications with Angular", CoverURL: getCoverURL("angular-in-action.jpg")},
		{ID: 35, Title: "Docker Deep Dive", Author: "Nigel Poulton", Category: "DevOps", Desc: "Zero to Docker in a single book", CoverURL: getCoverURL("docker-deep-dive.jpg")},
		{ID: 36, Title: "Kubernetes in Action", Author: "Marko Luksa", Category: "DevOps", Desc: "Learn Kubernetes from a developer perspective", CoverURL: getCoverURL("kubernetes-in-action.jpg")},
		{ID: 37, Title: "The DevOps Handbook", Author: "Gene Kim", Category: "DevOps", Desc: "How to create world-class agility, reliability, and security", CoverURL: getCoverURL("the-devops-handbook.jpg")},
		{ID: 38, Title: "Site Reliability Engineering", Author: "Betsy Beyer", Category: "DevOps", Desc: "How Google runs production systems", CoverURL: getCoverURL("site-reliability-engineering.jpg")},
		{ID: 39, Title: "Continuous Delivery", Author: "Jez Humble", Category: "DevOps", Desc: "Reliable software releases through automation", CoverURL: getCoverURL("continuous-delivery.jpg")},
		{ID: 40, Title: "Database Design for Mere Mortals", Author: "Michael Hernandez", Category: "Database", Desc: "A hands-on guide to relational database design", CoverURL: getCoverURL("database-design-for-mere-mortals.jpg")},
		{ID: 41, Title: "SQL Performance Explained", Author: "Markus Winand", Category: "Database", Desc: "Everything developers need to know about SQL performance", CoverURL: getCoverURL("sql-performance-explained.jpg")},
		{ID: 42, Title: "MongoDB: The Definitive Guide", Author: "Shannon Bradshaw", Category: "Database", Desc: "Powerful and scalable data storage", CoverURL: getCoverURL("mongodb-the-definitive-guide.jpg")},
		{ID: 43, Title: "Redis in Action", Author: "Josiah Carlson", Category: "Database", Desc: "Learn Redis through practical examples", CoverURL: getCoverURL("redis-in-action.jpg")},
		{ID: 44, Title: "Machine Learning Yearning", Author: "Andrew Ng", Category: "Machine Learning", Desc: "Technical strategy for AI engineers", CoverURL: getCoverURL("machine-learning-yearning.jpg")},
		{ID: 45, Title: "Hands-On Machine Learning", Author: "Aurélien Géron", Category: "Machine Learning", Desc: "With Scikit-Learn, Keras, and TensorFlow", CoverURL: getCoverURL("hands-on-machine-learning.jpg")},
		{ID: 46, Title: "Deep Learning", Author: "Ian Goodfellow", Category: "Machine Learning", Desc: "Comprehensive introduction to deep learning", CoverURL: getCoverURL("deep-learning.jpg")},
		{ID: 47, Title: "Pattern Recognition and Machine Learning", Author: "Christopher Bishop", Category: "Machine Learning", Desc: "A comprehensive treatment of the field", CoverURL: getCoverURL("pattern-recognition-and-machine-learning.jpg")},
		{ID: 48, Title: "Artificial Intelligence: A Modern Approach", Author: "Stuart Russell", Category: "Artificial Intelligence", Desc: "The definitive AI textbook", CoverURL: getCoverURL("artificial-intelligence-a-modern-approach.jpg")},
		{ID: 49, Title: "Designing Data-Intensive Applications", Author: "Martin Kleppmann", Category: "System Design", Desc: "The big ideas behind reliable, scalable systems", CoverURL: getCoverURL("designing-data-intensive-applications.jpg")},
		{ID: 50, Title: "System Design Interview", Author: "Alex Xu", Category: "System Design", Desc: "An insider's guide to system design", CoverURL: getCoverURL("system-design-interview.jpg")},
	}
	
	return books
}

func (ds *dataSource) PopulateData() error {
	err := ds.db.AutoMigrate(&mockapi.Book{})
	if err != nil {
		panic("failed to migrate database")
	}
	return ds.db.Transaction(func(tx *gorm.DB) error {
		var count int64
		if err := tx.Model(&mockapi.Book{}).Count(&count).Error; err != nil {
			return err
		}

		if count > 0 {
			return nil
		}

		books := getInitialBooks()

		if err := tx.Create(&books).Error; err != nil {
			return err
		}

		return nil
	})
}

func (ds *dataSource) GetBookByID(id string) (mockapi.Book, error) {
	var book mockapi.Book
	if err := ds.db.First(&book, "id = ?", id).Error; err != nil {
		return mockapi.Book{}, err
	}
	return book, nil
}

func (ds *dataSource) GetBooks(page int, pageSize int, search string) ([]mockapi.Book, error) {
	var books []mockapi.Book
	offset := (page - 1) * pageSize
	
	if search != "" {
		searchPattern := "%" + search + "%"
		if err := ds.db.Offset(offset).Limit(pageSize).Where("title LIKE ? OR author LIKE ?", searchPattern, searchPattern).Find(&books).Error; err != nil {
			return nil, err
		}
		return books, nil
	}

	if err := ds.db.Offset(offset).Limit(pageSize).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (ds *dataSource) GetBooksCount(search string) (int64, error) {
	var count int64
	if search != "" {
		searchPattern := "%" + search + "%"
		if err := ds.db.Model(&mockapi.Book{}).Where("title LIKE ? OR author LIKE ?", searchPattern, searchPattern).Count(&count).Error; err != nil {
			return 0, err
		}
		return count, nil
	}

	if err := ds.db.Model(&mockapi.Book{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}