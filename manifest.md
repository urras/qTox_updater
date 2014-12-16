The manifest format is fairly simple.

``[(magic number){0xCA, 0xFE, 0xBA, 0xBE}][(size of next item) uint8][File_data.Name string][(size of next item) uint8][File_data.Hash (sha1) string][File_data.Buildnumber uint8]``

Note: SHA1 hashes are the same size always and I don't see build numbers going out of control, but a file with a name over 255 characters isn't impossible.

After the magic number, the next byte tell us the size of the next item, going ``[size, name, size, hash, build number]``. We continue to iterate through the manifest until we hit a ``0x127`` (esc)

If this format is invalid in any way or a hash doesn't add up the update is stopped and qTox is not touched. If a build number is higher than the current internal build number said file is copied and its hash is validated, reverting it if it doesn't add up to the hash stored in the File_data object.