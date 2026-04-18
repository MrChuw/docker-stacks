#!/bin/bash

set -e  # Exit immediately if a command fails

EXTENSIONS_DIR="./config/www/freshrss/extensions"
FRESHRSS_EXT_REPO="https://github.com/FreshRSS/Extensions"
FRESHRSS_EXT_DIR="./freshrss_extensions"

# Cloning or updating the freshrss_extensions repository
if [ -d "$FRESHRSS_EXT_DIR/.git" ]; then
    echo "Updating freshrss_extensions..."
    git -C "$FRESHRSS_EXT_DIR" pull
else
    echo "Cloning freshrss_extensions..."
    git clone "$FRESHRSS_EXT_REPO" "$FRESHRSS_EXT_DIR"
fi

# Syncing selected extensions
EXTENSIONS_TO_SYNC=(
    "xExtension-ColorfulList"
    "xExtension-ReadingTime"
    "xExtension-showFeedID"
    "xExtension-TitleWrap"
    "xExtension-YouTube"
)

for ext in "${EXTENSIONS_TO_SYNC[@]}"; do
    echo "Syncing $ext..."
    rsync -av --delete "$FRESHRSS_EXT_DIR/$ext/" "$EXTENSIONS_DIR/$ext/"
done

# Cloning or updating external repositories
EXT_REPOS=(
    "https://github.com/printfuck/xExtension-Readable"
    "https://github.com/aledeg/xExtension-LatexSupport"
    "https://github.com/aledeg/xExtension-DateFormat"
    "https://github.com/mgnsk/FreshRSS-AutoTTL"
    "https://framagit.org/nicofrand/xextension-togglablemenu"
    "https://github.com/ravenscroftj/freshrss-flaresolverr-extension"
)

for repo in "${EXT_REPOS[@]}"; do
    ext_name=$(basename "$repo")
    ext_path="$EXTENSIONS_DIR/$ext_name"

    if [ -d "$ext_path/.git" ]; then
        echo "Updating $ext_name..."
        git -C "$ext_path" pull
    else
        echo "Cloning $ext_name..."
        git clone "$repo" "$ext_path"
    fi

done

echo "Extensions updated successfully!"