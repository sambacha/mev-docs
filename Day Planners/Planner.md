```markdown
## Day Planner

This is my plan for the day broken into 3 main sections:
1. Morning Prep
2. Reading
3. Afternoon Review

### Morning Prep

This is where I get ready for work and do my usual prep.

- [ ] 09:30 Setup for work
- [ ] 09:45 Review notes from yesterday
- [ ] 10:30 Create new notes for #article review
- [ ] 11:30 BREAK

### Reading

A section of the day dedicated to reading:

1. Articles.
2. Book chapters assigned for the day.
3. Re-reading past notes.
   
- [ ] 12:00 Reading
  - [ ] Article 1
  - [ ] Article 2
  - [ ] Article notes review
- [ ] 12:25 BREAK
- [ ] 12:30 Reading
- [ ] 14:00 BREAK

### Afternoon Review

I use this time to review what I have done earlier in the day and complete any tasks to prepare for the next day.

- [ ] 15:00 Review notes and update daily note [[20201103]]
- [ ] 15:45 Walk
- [ ] 16:30 Reading
- [ ] 17:20 Prep for tomorrow's meetings
- [ ] 18:00 END
```

This is also provided as a file in [day-planner-example.md](https://github.com/lynchjames/obsidian-day-planner/blob/main/examples/day-planner-example.md).

The Day Planner heading and `---` rule are used to identify the extent of the Day Planner. A heading must be used but can be `#`, `##`, `###` or `####`.

The format of the task list items is important as this is what is used to calculate the times of each task and the intervals between tasks. The format used should be:

`- [ ] HH:mm Task text`

**24 hour times should be used.**

Nested checklist items or bullets are now also supported to capture sub-tasks of a timed task. Timed tasks must be at the top level of the checkbox list.

`BREAK` and `END` are keywords that define breaks and the end to the time tracking for the tasks. They are not case sensitive so `break` and `end` can also be used.

`END` is used as an item with a time to give an accurate time interval for the last task, _"Prep for tomorrow's meetings"_ at 17:00 in this example.

The note will update automatically: tasks in the past will be checked and marked as complete.

Using the example above, at 14:30 the note would have automatically updated to:

```markdown
## Day Planner

This is my plan for the day broken into 3 main sections:
1. Morning Prep
2. Reading
3. Afternoon Review

### Morning Prep

This is where I get ready for work and do my usual prep.

- [x] 09:30 Setup for work
- [x] 09:45 Review notes from yesterday
- [x] 10:30 Create new notes for #article review
- [x] 11:30 BREAK

### Reading

A section of the day dedicated to reading:

1. Articles.
2. Book chapters assigned for the day.
3. Re-reading past notes.
   
- [x] 12:00 Reading
  - [ ] Article 1
  - [ ] Article 2
  - [ ] Article notes review
- [x] 12:25 BREAK
- [x] 12:30 Reading
- [ ] 14:00 BREAK

### Afternoon Review

I use this time to review what I have done earlier in the day and complete any tasks to prepare for the next day.

- [ ] 15:00 Review notes and update daily note [[20201103]]
- [ ] 15:45 Walk
- [ ] 16:30 Reading
- [ ] 17:20 Prep for tomorrow's meetings
- [ ] 18:00 END
```