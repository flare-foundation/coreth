#!/bin/bash
# Requires wkhtmltopdf and aha tools
# Install using: sudo apt install wkhtmltopdf aha
git diff --color upstream-v0.6.2-rc.0 | aha > CHANGES.html && wkhtmltopdf CHANGES.html CHANGES.pdf && rm CHANGES.html