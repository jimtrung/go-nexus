CREATE OR REPLACE FUNCTION auto_accept_mutual_friend_request() 
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status = 'pending' AND 
       EXISTS (SELECT 1 FROM friends WHERE sender_id = NEW.receiver_id AND receiver_id = NEW.sender_id AND status = 'pending') THEN
        
        UPDATE friends 
        SET status = 'accepted' 
        WHERE sender_id = NEW.receiver_id AND receiver_id = NEW.sender_id;
        
        NEW.status := 'accepted';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Instead of AFTER INSERT, use BEFORE INSERT so we can modify NEW
CREATE TRIGGER auto_accept_mutual_friend_request_trigger
BEFORE INSERT ON friends
FOR EACH ROW EXECUTE FUNCTION auto_accept_mutual_friend_request();
