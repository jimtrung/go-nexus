CREATE TRIGGER friends_update_timestamp
BEFORE UPDATE ON friends
FOR EACH ROW
EXECUTE FUNCTION update_timestamp_trigger();