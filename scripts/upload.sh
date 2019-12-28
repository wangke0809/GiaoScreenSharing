#!/bin/sh

uploadFile()
{
    FILE=$1
    BASENAME="$(basename "${FILE}")"
    uploadedTo=`wget --method PUT --body-file=$FILE "https://transfer.sh/$BASENAME" -O - -nv`
    echo "$BASENAME uploaded to url: $uploadedTo"
}

zip -r ScreenSharing.zip *

uploadFile "ScreenSharing.zip"