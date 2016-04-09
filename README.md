# cs496_assignment2

For this assignment you will make a dynamic webpage that can store, edit and display data.

The requirements for this assignment are as follows:

You should host the page using a cloud based provider (public_html at OSU will not suffice)
Users should be able to input data using different sorts of form controls. There should be at least 5 different kinds of input controls corresponding to 5 different properties. (See 4.10.5.1 (Links to an external site.) for a list of possible inputs)
One such control should be a check-box
Another should be a radio button
The others are up to you
The lectures show uploading an image. I would recommend not doing this as it is more difficult than other kinds of data, but I wanted to make sure how to illustrate it in the lecture as it is a common use case.
The user should be able to add an arbitrary number of new records.
Users should be able to view the data for any record they previously entered in a state such that it cannot be changed. In other words a page to just 'view' the data.
Users should be able to enter an edit mode for any record (either on the same page or loaded on a different page) allowing them to change the data
This edit mode should prepopulate all needed data. In other words, if I edit an item, do nothing but click 'save changes' nothing should happen to the item.
Upon saving the object should have the same id but have updated information (if this is not possible on your platform, document why not)
Deliverables:

You should submit a PDF containing the following
A URL to the live site
A description of the site and what data can be saved (a paragraph or 2)
A test plan to make sure it works as intended (I would guess about 25 test cases would cover something like this)
This plan can be manual tests, no need to automate anything
It should cover basic edge cases like saving empty items or duplicate items
And example test might be: "Input: Save user with no last name. Expected result: user is not saved, message indicating last name is required is displayed."
The results of executing your test plan (Successes should be a simple success, failures should be a few sentences to a couple paragraphs)
Did all of your tests pass?
If tests failed, detail how and try to explain why
Discuss how you used the idea of templating to display data.(1 or 2 paragraphs)
Discuss if you had it to do over again, what changes you might make. (2 or 3 paragraphs)
Notes about future projects:

Looking to the future, you should pick a subject that is sufficiently complex and interests you as you may be able to leverage some of this early work into a final project. To that end, think about a piece of software that would have to keep track of at least 3 different kinds of objects that have several properties each.(including lists of things, for example a user might have a list of favorite programming languages in a programming job board)

Additionally it would be good if there was a use for the data (or interaction with the data) outside of your user interface as an upcoming assignment will require you to make an API rather than a web page. For example, a company might want to design a custom HR application specific to their needs that pulled information from the previously mentioned job board. They would use an API to do such a task.

Rubric
1) At least 5 different input controls (For e.g: "type=text" for Name input, "type=url" for Website input, "type=email" for Email input, "type=password" for account Password input, "type=number" for Age input). 
ONE such input will be a check-box. ONE such input will be a radio button

2.5 pts

2) Saving - Data should save without issues - No error messages should show up - No syntax elements should be visible to the client anywhere on the website - Proper messages (pop-ups, alerts, etc) should relay information at all times to the user Viewing - Contents of a view page should be unalterable (View state and Save/Edit states, are not/should not, be the same)

2.5 pts

3) - Editing - Correctly pre-populates all needed data - Data should be changed without needing to assign/create a new ID - Correctly handles all scenarios (saving without editing anything, change only one item, not providing required inputs, etc.) - Should not forcibly make all fields blank when entering/exiting edit state

1.5 pts

4) write up

0 pts

5) -- Test plan. Should have AT LEAST 10 test cases ( or a good reason for not doing so ) Should cover all basic, edge, blank. Should have test plan results. Should intelligibly explain failures.

2 pts

6) Reflection Should have a paragraph or 2 describing changes and why

1.5 pts

