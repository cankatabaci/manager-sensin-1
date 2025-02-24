package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	LongName     string             `bson:"long_name,omitempty"`
	Name         string             `bson:"short_name,omitempty"`
	Positions    string             `bson:"player_positions,omitempty"`
	ClubPosition string             `bson:"club_position,omitempty"`
	Club         string             `bson:"club_name,omitempty"`
	League       string             `bson:"league_name,omitempty"`
	Nationality  string             `bson:"nationality_name,omitempty"`
	Age          int                `bson:"age,omitempty"`
	Overall      int                `bson:"overall,omitempty"`
	Potential    int                `bson:"potential,omitempty"`
	Pace         interface{}        `bson:"pace,omitempty"`
	Passing      interface{}        `bson:"passing,omitempty"`
	Physic       interface{}        `bson:"physic,omitempty"`
	Shooting     interface{}        `bson:"shooting,omitempty"`
	Dribbling    interface{}        `bson:"dribbling,omitempty"`
	Defending    interface{}        `bson:"defending,omitempty"`
	FaceUrl      string             `bson:"player_face_url,omitempty"`
	ClubLogo     string             `bson:"club_logo_url,omitempty"`
	NationFlag   string             `bson:"nation_flag_url,omitempty"`
	WF           int                `bson:"weak_foot,omitempty"`
	SM           int                `bson:"skill_moves,omitempty"`
	WorkRate     string             `bson:"work_rate,omitempty"`
	Foot         string             `bson:"preferred_foot,omitempty"`
	Hidden       bool               `bson:"hidden,omitempty"`
}

type Manager struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Email   string             `json:"email,omitempty" bson:"email,omitempty"`
	UserID  string             `json:"userID,omitempty" bson:"userID,omitempty"`
	Points  int                `bson:"points,omitempty"`
	Players []Player           `bson:"players,omitempty"`
	Results []Result           `bson:"results,omitempty"`
}

type Result struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Season      string             `json:"season,omitempty" bson:"season,omitempty"`
	Home        string             `json:"home,omitempty" bson:"home,omitempty"`
	Away        string             `json:"away,omitempty" bson:"away,omitempty"`
	SeasonType  string             `json:"seasonType,omitempty" bson:"seasonType,omitempty"`
	SeasonTitle string             `json:"seasonTitle,omitempty" bson:"seasonTitle,omitempty"`
	HomeManager string             `json:"homeManager,omitempty" bson:"homeManager,omitempty"`
	AwayManager string             `json:"awayManager,omitempty" bson:"awayManager,omitempty"`
	Score       []int              `json:"score,omitempty" bson:"score,omitempty"`
	HomeScorers []Scorer           `json:"homescorers,omitempty" bson:"homescorers,omitempty"`
	AwayScorers []Scorer           `json:"awayscorers,omitempty" bson:"awayscorers,omitempty"`
}

type Season struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Type     string             `json:"type,omitempty" bson:"type,omitempty"`
	Title    string             `json:"title,omitempty" bson:"title,omitempty"`
	IsActive bool               `json:"isActive" bson:"isActive"`
	Results  []Result           `bson:"results,omitempty"`
}

type Scorer struct {
	Player Player `json:"player,omitempty"`
	Count  int    `json:"count,omitempty"`
}

type Insert struct {
	InsertedID primitive.ObjectID
}

type Standing struct {
	Manager string   `json:"manager"`
	Points  int      `json:"points"`
	Played  int      `json:"played"`
	Won     int      `json:"won"`
	Draw    int      `json:"draw"`
	Lost    int      `json:"lost"`
	GF      int      `json:"gf"`
	GA      int      `json:"ga"`
	GD      int      `json:"gd"`
	Form    []string `json:"form"`
}
type Stats struct {
	Manager string `json:"manager"`
	Player  string `json:"player"`
	Count   int    `json:"count"`
	FaceUrl string `json:"faceUrl"`
}

//manager-logic
func (m *Manager) playerExist(playerID primitive.ObjectID) (bool, int) {
	found := false
	foundIndex := 0
	for i, player := range m.Players {
		if player.ID == playerID {
			found = true
			foundIndex = i
			break
		}
	}
	return found, foundIndex
}
func (m *Manager) AddPlayer(p Player) {
	found, _ := m.playerExist(p.ID)
	if !found {
		m.Players = append(m.Players, p)
	}
}
func (m *Manager) DeletePlayer(p Player) {
	found, index := m.playerExist(p.ID)
	if found {
		m.Players = append(m.Players[:index], m.Players[index+1:]...)
	}
}

func (m *Manager) ManagePoint(point, pointType int) {
	if pointType == 0 {
		m.Points -= point
	} else {
		m.Points += point
	}
}

func (m *Manager) AddResult(r Result) {
	m.Results = append(m.Results, r)
}

//season-logic
func (s *Season) ChangeStatus(isActive bool) {
	s.IsActive = isActive
}
func (s *Season) AddResult(r Result) {
	s.Results = append(s.Results, r)
}
func (s *Standing) Set(standing Standing) {
	s.Manager = standing.Manager
	s.Points += standing.Points
	s.Played += standing.Played
	s.Won += standing.Won
	s.Draw += standing.Draw
	s.Lost += standing.Lost
	s.GF += standing.GF
	s.GA += standing.GA
	s.GD += standing.GD
	s.Form = append(s.Form, standing.Form...)
}
