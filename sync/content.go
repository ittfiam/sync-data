package sync


const MYSQLREADER = 1

const MYSQLWRITER = -1

const MONGOREADER = 2

const MONGOWRITER = -2


type Speed struct {
	Channel int `json:"channel"`
}

func newSpeed() *Speed{
	return &Speed{
		Channel:1,
	}
}

type Setting struct{
	Sped *Speed `json:"speed"`
}

func newSetting() *Setting{

	return &Setting{
		Sped:newSpeed(),
	}
}


type Content struct {
	Reader interface{} `json:"reader"`
	Writer interface{} `json:"writer"`
}


func newContent(reader interface{},writer interface{}) *Content{

	return &Content{
		Reader:reader,
		Writer:writer,
	}
}


type Jobs struct {

	Content []*Content `json:"plugin"`

	Sett *Setting `json:"setting"`

}


func newJobs() *Jobs{
	return &Jobs{
		Content:make([]*Content,0),
		Sett:newSetting(),
	}
}


type Work struct {
	Job *Jobs `json:"job"`
}



func NewWorker(reader interface{},writer interface{}) *Work{

	c := newContent(reader,writer)
	job := newJobs()
	job.Content = append(job.Content, c)

	return &Work{
		Job:job,
	}
}
