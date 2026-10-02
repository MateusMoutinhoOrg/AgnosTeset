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
    os.system("agnos start --project-name testebackoffice --module github.com/MateusMoutinhoOrg/AgnosTeset")
    os.system("agnos backoffice-init")
   



if __name__ == '__main__':
    main()
