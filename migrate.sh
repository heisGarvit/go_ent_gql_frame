source .env
echo "Migrating database..."
date
atlas migrate diff migration --dir "file://atlas" --to "ent://ent/schema" --dev-url "$BLANK_DB_URI"
ls -lrt atlas
git add atlas