This repo is used for Git training.

Instructions:
- Clone the repository
- Create a feature branch under your name
- Complete the assigned task
- Push the branch and raise a Pull Request


Cloning - 

Name: Vishesh Chandok  
Date: 07-01-2026  
Interested Project: Data Analytics / Machine Learning based projects


# Git Undo Practice: Reset vs Revert

This project was created to demonstrate how to safely handle mistakes in a Git repository. I intentionally committed an error to practice two different recovery methods: `git reset --soft` and `git revert`.

## The "Mistake"
I committed a file where I divided a integer with 0.

---

## Method Used: git revert

For the final submission, I chose to use `git revert`.

### Why I chose Revert over Reset:
1. **Preserving Context:** Reverting allows me to keep the record of the mistake. In a learning environment, this is helpful to show the "before and after" of how the bug was handled.
2. **Public History:** Since this repo is intended for Openverse/GitHub, using `revert` ensures that anyone who has already pulled the code won't experience a "broken" history (which happens if you delete commits using reset).

---

## Command Breakdown

### 1. git reset --soft
* **The Logic:** This "undoes" the commit but leaves the faulty code in my staging area. 
* **Scenario:** I would use this if I realized I forgot to add a `rescue` block *before* I pushed my code to the main branch.

### 2. git revert
* **The Logic:** This creates a new "inverse" commit. If the original commit added `x / 0`, the revert commit effectively removes it while moving the project forward.
* **Scenario:** Best for fixing the crash after the code has already been shared with the team.

---
