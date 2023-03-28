# golang-tmux-poc

This repository contains the proof of concept for made for [kite](https://github.com/openshift/pagerduty-short-circuiter) for multiple-terminal windows similar to tmux. 

# Prerequisites
Go installed in the local computer with version >= 1.18.

[tterm](https://git.sr.ht/~rockorager/tcell-term) Repository is used to implement the tmux approach.

# Usage

The app uses the following navigation:

* Ctrl N -> Next Slide
* Ctrl P -> Previous Slide
* Ctrl A -> Add Slide
* Ctrl E -> Exit Slide

Example usage:

```
go run  .
```

# Contributing
If you would like to contribute to the project, please fork the repository and make your changes. 

Once you have finished your changes, please submit a pull request for review.







