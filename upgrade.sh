source .env
echo "Upgrading database..."
date
/usr/local/bin/atlas migrate apply --dir file://migrations --url "$DB_URI"