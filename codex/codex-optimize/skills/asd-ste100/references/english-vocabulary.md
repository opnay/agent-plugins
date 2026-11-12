# English Vocabulary And Usage

Read this reference when drafting or reviewing English prose. It covers English-specific guidance; do not transfer its vocabulary, parts of speech, spelling, or grammar restrictions to Korean or other languages.

This is a practical summary with selected examples from the official STEMG FAQ. It is not the complete ASD-STE100 dictionary or rule set. For formal compliance, also read [Standard scope and evidence](standard-scope.md).

## General Vocabulary

- Check a general word's approved meaning and part of speech in the applicable edition's dictionary. Familiar spelling alone does not establish approval for every sense or grammatical role.
- Use one consistent term for a concept. Prefer an approved alternative only when it preserves the source meaning and fits the sentence.
- STE approved meanings and spelling are based on American English and Merriam-Webster. Do not change protected literals, product names, or quoted text to enforce spelling.
- If dictionary evidence is unavailable, distinguish a clarity suggestion from a verified vocabulary correction. Do not reconstruct the full dictionary from memory or assume unlisted words are approved.

Selected official FAQ examples:

| Word or choice | Restriction illustrated by the FAQ | Editing implication |
| --- | --- | --- |
| `start` | STE selects this synonym rather than `begin`, `commence`, `initiate`, or `originate`. | Consider `start` when the intended meaning matches; preserve literal commands and UI labels. |
| `fall` | The approved verb sense concerns movement down by gravity, not a decrease. | Do not describe a falling metric with this sense; check the suitable dictionary entry. |
| `about` | The approved sense is “concerned with,” not “approximately” or “around.” | Identify the intended meaning before selecting a different word. |
| `check` | Approved as a noun, not a verb. | The FAQ contrasts `do a check` with `check the lights`; inspect the intended action before rewriting. |

These examples are not a global search-and-replace list. They do not approve replacement words in every context.

## Technical Nouns And Technical Verbs

STE permits subject-specific technical nouns and technical verbs under its terminology rules. They are not an unrestricted escape from the general dictionary.

- Establish the term from official documentation, engineering drawings, an approved company glossary, or a terminology database.
- Record the intended concept and scope when needed. Use the exact established term consistently.
- Distinguish a technical term from an ordinary word used informally. Do not relabel an unapproved general word as a technical term merely to accept it.
- Preserve necessary software identifiers, commands, paths, and UI labels as literals. Their presence does not prove that the surrounding prose complies.

## English Grammar And Sentence Type

- Procedures use direct instructions in the imperative and do not use passive sentences.
- Descriptions use active voice. The official FAQ permits passive voice when the actor is unknown; do not invent an actor to remove it.
- Put a condition before the main clause when readers must know it before doing the work step.
- `-ing` forms are generally restricted because their grammatical function can be ambiguous. The restriction has exceptions for technical nouns and approved words such as `opening`, `remaining`, `something`, and `during`. Check the role and applicable rule; do not ban a suffix mechanically.
- For exact sentence-length limits, word-count conventions, permitted constructions, and exceptions, consult the selected edition. This summary does not establish those checks.

## Meaning-Preserving Examples

The following are local clarity examples, not assertions that every word has passed a full dictionary review.

### Split Sequential Actions

Source:

> Close the application and then disconnect the cable.

Rewrite:

> 1. Close the application.
> 2. Disconnect the cable.

Keep the sequence. Do not change the actions into alternatives or perform them concurrently.

### Put A Known Condition First

Source:

> Restart the service if the status is `FAILED`.

Rewrite:

> If the status is `FAILED`, restart the service.

Preserve the exact status and the conditional scope.

### Retain Missing Criteria As A Gap

Source:

> Replace the filter if necessary.

The replacement criterion is not supplied. Flag the missing criterion instead of inventing a damage threshold or changing this into an unconditional replacement instruction.

## Source

- [Official STEMG FAQ](https://www.asd-ste100.org/STE_faq.html): general vocabulary, examples, terminology, conditions, voice, and `-ing` restrictions.
- Public guidance checked on 2026-10-04. The site identifies Issue 9, dated January 15, 2025. The complete standard and dictionary were not reviewed for this summary; use the required edition for formal checks.
