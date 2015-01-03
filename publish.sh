#!/bin/bash

[[ -d fatt.club/public ]] && rsync -av fatt.club/public/* fatt.club:web/ 2>&1 | xmessage -file -
