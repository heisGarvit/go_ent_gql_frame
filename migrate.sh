source .env
echo "Migrating database..."
date
/usr/local/bin/atlas migrate diff migration --dir "file://migrations" --to "ent://src/models/ent/schema" --dev-url "$BLANK_DB_URI"
ls -lrt migrations
git add migrations