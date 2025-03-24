CREATE OR REPLACE FUNCTION auto_accept_mutual_friend_request() RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status = 'pending' THEN
        UPDATE friends SET status = 'accepted'
        WHERE sender_id = NEW.receiver_id AND receiver_id = NEW.sender_id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER auto_accept_mutual_friend_request_trigger
AFTER INSERT ON friends
FOR EACH ROW EXECUTE FUNCTION auto_accept_mutual_friend_request();