source .env
echo "Upgrading database..."
date
atlas migrate apply --dir file://atlas --url "$DB_URI"