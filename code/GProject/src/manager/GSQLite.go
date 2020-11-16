//===============================================
package manager
//===============================================
import "fmt"
import "sync"
import "database/sql"
import _ "github.com/mattn/go-sqlite3"
//===============================================
// struct
//===============================================
type GSQLiteO struct {

}
//===============================================
// constructor
//===============================================
var m_GSQLiteLock = &sync.Mutex{}
var m_GSQLiteO *GSQLiteO
//===============================================
func newGSQLite() *GSQLiteO {
    lObj := &GSQLiteO{}
    return lObj
}
//===============================================
func GSQLite() *GSQLiteO {
    if m_GSQLiteO == nil {
        m_GSQLiteLock.Lock()
        defer m_GSQLiteLock.Unlock()
        if m_GSQLiteO == nil {
            m_GSQLiteO = newGSQLite()
        }
    }
    return m_GSQLiteO
}
//===============================================
// private
//===============================================
func (obj *GSQLiteO) open() *sql.DB {
    lSqlite := GManager().mgr.sqlite
    lDb, _ :=  sql.Open("sqlite3", lSqlite.db_path)
    return lDb
}
//===============================================
// public
//===============================================
func (obj *GSQLiteO) QueryShow(sqlQuery string) {
    lDb := obj.open()
    lRows, _ := lDb.Query(sqlQuery)
    lColNames, _ := lRows.Columns()
    lColsId := make([]interface{}, len(lColNames))
    lCols := make([]string, len(lColNames))
    for i := range lColsId {
        lColsId[i] = &lCols[i]
    }
    for lRows.Next() { 
        lRows.Scan(lColsId...)
        for i := 0; i < len(lCols); i++ {
            if i != 0 {fmt.Printf(" | ")}
            fmt.Printf("%s", lCols[i])
        }
        fmt.Printf("\n")
    }
    lRows.Close()
    lDb.Close()
}
//===============================================
func (obj *GSQLiteO) QueryCol(sqlQuery string) []string {
    lDb := obj.open()
    lRows, _ := lDb.Query(sqlQuery)
    lColNames, _ := lRows.Columns()
    lColsId := make([]interface{}, len(lColNames))
    lCols := make([]string, len(lColNames))
    var lData []string
    for i := range lColsId {
        lColsId[i] = &lCols[i]
    }
    for lRows.Next() { 
        lRows.Scan(lColsId...)
        for i := 0; i < len(lCols); i++ {
            lData = append(lData, lCols[i])
            break
        }
    }
    lRows.Close()
    lDb.Close()
    return lData
}
//===============================================
