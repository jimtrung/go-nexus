CREATE OR REPLACE FUNCTION accept_mutual_friend_requests()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM friends
        WHERE sender_id = NEW.receiver_id
          AND receiver_id = NEW.sender_id
          AND status = 'pending'
    ) THEN
        UPDATE friends
        SET status = 'accepted'
        WHERE (sender_id = NEW.sender_id AND receiver_id = NEW.receiver_id)
           OR (sender_id = NEW.receiver_id AND receiver_id = NEW.sender_id);
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
