package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/danilocordeirodev/social-go/internal/store"
)

var usernames = []string{
	"alice", "bob", "arthur", "charlie", "diana", "eve", "frank", "grace", "hank", "isla",
	"jack",
	"kira",
	"leo",
	"mona",
	"nina",
	"oliver",
	"penny",
	"quinn",
	"ryan",
	"sara",
	"tom",
	"uma",
	"viktor",
	"willa",
	"xander",
	"yara",
	"zane",
	"amber",
	"brad",
	"clara",
	"david",
	"ella",
	"finn",
	"gwen",
	"hugo",
	"ivy",
	"jude",
	"kara",
	"luke",
	"mia",
	"noah",
	"ocean",
	"paul",
	"quincy",
	"rose",
	"simon",
	"tina",
	"uri",
	"vivi",
}

var titles = []string{
	"10 Tips for Boosting Your Productivity",
	"The Ultimate Guide to Mindful Living",
	"Exploring the Benefits of Remote Work",
	"How to Create a Sustainable Morning Routine",
	"Top 5 Must-Read Books for Personal Growth",
	"The Art of Minimalism: Less is More",
	"Traveling on a Budget: Tips and Tricks",
	"The Power of Habit: How to Build Lasting Change",
	"A Beginner's Guide to Plant-Based Eating",
	"The Science Behind Effective Goal Setting",
	"Unleashing Your Creativity: Daily Practices",
	"Navigating Life Changes: Strategies for Success",
	"The Importance of Self-Care in a Busy World",
	"Digital Detox: Reclaiming Your Time and Focus",
	"How to Cultivate Resilience in Tough Times",
	"Exploring New Hobbies: Finding Your Passion",
	"Mindfulness Techniques for Stress Relief",
	"Building Stronger Relationships in a Digital Age",
	"The Future of Work: Trends to Watch",
	"Healthy Meal Prep Ideas for Busy Weeks",
	"Finding Balance: Work, Life, and Everything In Between",
}

var contents = []string{
	"Explore the benefits of meditation and how to get started in just five minutes a day.",
	"A deep dive into the latest trends in sustainable fashion and how to make eco-friendly choices.",
	"Tips for creating a productive home office space that inspires creativity and focus.",
	"Interviews with local entrepreneurs on how they turned their passions into successful businesses.",
	"An exploration of the science behind sleep and tips for improving your nightly rest.",
	"A guide to the best online resources for learning new skills during the pandemic.",
	"Personal anecdotes about overcoming challenges and what they taught me about resilience.",
	"A round-up of the best apps for managing your time and boosting productivity.",
	"Fun and easy DIY projects to brighten up your living space without breaking the bank.",
	"How to develop a gratitude practice that can improve your mental well-being.",
	"A travel journal highlighting off-the-beaten-path destinations worth exploring.",
	"Explaining the concept of growth mindset and how it can transform your approach to challenges.",
	"Recipes for quick, healthy meals that can be prepared in under 30 minutes.",
	"Tips for cultivating a garden at home, even in small spaces.",
	"An analysis of the impact of social media on mental health, backed by recent studies.",
	"Step-by-step instructions for hosting a virtual event that engages participants.",
	"A list of must-visit local parks and outdoor spaces for a weekend getaway.",
	"Insights into how to create a personal budget and stick to it effectively.",
	"Ideas for incorporating mindfulness into your daily routine for better mental clarity.",
	"A guide to understanding and embracing different learning styles.",
	"Personal reflections on the importance of community service and giving back.",
}
var tags = []string{
	"productivity",
	"mindfulness",
	"sustainability",
	"personal development",
	"health",
	"travel",
	"self-care",
	"entrepreneurship",
	"wellness",
	"creativity",
	"budgeting",
	"DIY",
	"mental health",
	"food",
	"lifestyle",
	"technology",
	"outdoors",
	"inspiration",
	"community",
	"growth mindset",
}

var comments = []string{
	"Great insights! I really enjoyed this post.",
	"Thanks for sharing your thoughts. Very helpful!",
	"I love the tips you provided. Can’t wait to try them!",
	"This was an eye-opener for me. Appreciate the information!",
	"Such an inspiring read! Keep up the good work.",
	"I completely agree with your perspective. Well said!",
	"This article has motivated me to make some changes in my life.",
	"I found this very informative. Thank you for writing it!",
	"Interesting points! I hadn’t considered that before.",
	"Your writing style is engaging and easy to follow.",
	"Thanks for the practical advice. I’ll definitely implement it.",
	"I appreciate the personal anecdotes. They made it relatable.",
	"This is exactly what I needed to hear today. Thank you!",
	"You covered so many important aspects. Well done!",
	"I love how you break down complex topics. Very helpful!",
	"This post sparked a lot of ideas for me. Thank you!",
	"I always look forward to your content. It’s so inspiring!",
	"Fantastic job! I learned a lot from this.",
	"You have a way with words that really resonates with me.",
	"Thanks for tackling such an important subject!",
	"I’m bookmarking this for future reference. Great post!",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(200, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123123",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
