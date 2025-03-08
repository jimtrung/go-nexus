CREATE TRIGGER trigger_accept_mutual_requests
AFTER INSERT ON friends
FOR EACH ROW
EXECUTE FUNCTION accept_mutual_friend_requests();
