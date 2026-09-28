import os
import shutil

def remove_all_files():
    repo_dir = os.path.dirname(os.path.abspath(__file__))
    script_name = os.path.basename(__file__)
    
    for item in os.listdir(repo_dir):
        # Ignora o próprio script
        if item == script_name:
            continue
        
        # Ignora a pasta do git (se houver) para manter o histórico
        if item == '.git':
            continue
            
        item_path = os.path.join(repo_dir, item)
        
        try:
            if os.path.isfile(item_path) or os.path.islink(item_path):
                os.unlink(item_path)
                print(f"Removido: {item}")
            elif os.path.isdir(item_path):
                shutil.rmtree(item_path)
                print(f"Removida (pasta): {item}")
        except Exception as e:
            print(f"Erro ao remover {item}: {e}")

def main():
    remove_all_files()
    ## runs a terminal command 
    os.system("agnos start --project-name teste --module github.com/MateusMoutinhoOrg/AgnosTeset")
    os.system("agnos front-init")
    os.system('agnos add-route adminmiddlware --pattern "/admin/{*rest}" --middleware  ')
    ## front route
    os.system('agnos add-route users --pattern "/admin/users" --category frontend --method GET -- response-type "text/html"')
    os.system('agnos add-route userconfig --pattern "/admin/user/{user:integer}" --category frontend --method GET -- response-type "text/html"')

    ## Api 
    os.system('agnos add-route apimiddleware --pattern "/api/admin/{*rest}" --middleware  ')
    os.system('agnos add-parameter token --route apimiddleware')
    os.system('agnos add-route apiadduser  --pattern "/api/admin/add-user" --category api --method POST --response-type "application/json" ')
    os.system('agnos add-route apieremoveuser  --pattern "/api/admin/remove-user" --category api --method POST --response-type "application/json"')
    os.system('agnos add-route apiupdateuser  --pattern "/api/admin/update-user" --category api --method POST --response-type "application/json"')
    os.system('agnos add-route apilistusers  --pattern "/api/admin/list-users" --category api --method GET --response-type "application/json"')



if __name__ == '__main__':
    main()
