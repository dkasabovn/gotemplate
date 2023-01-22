CURRENT=$(grep -oP '(?<=module\s{1})(\S{0,30})' go.mod)

echo "Currently Project is named $CURRENT"
echo "Input the new name:"

read NAME

echo "Changing package name to $NAME!"

sed -i "s/ProjectName = \"$CURRENT\"/ProjectName = \"$NAME\"/g" magefile.go
sed -i "s/module $CURRENT/module $NAME/g" go.mod

FILES=$(find app/ -type f -name "*.go")

for FILE in $FILES; do sed -i "s/$CURRENT\//$NAME\//g" $FILE ; done

echo "All done!"
